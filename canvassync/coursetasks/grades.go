package coursetasks

import (
	"fmt"
	"sort"
	"strings"
	"time"

	"github.com/zachdeibert/canvas-sync/canvas"
	"github.com/zachdeibert/canvas-sync/csvgen"
	"github.com/zachdeibert/canvas-sync/task"
)

type gradedAssignment struct {
	Name     string
	Due      time.Time
	ID       int
	Position int
	Score    float64
	MaxScore float64
	Dropped  bool
	Late     bool
	Missing  bool
	Excused  bool
	Graded   bool
	Section  *gradeSection
}

func (a gradedAssignment) csv(csv csvgen.CSV) {
	percentage := a.Score * 100 / a.MaxScore
	if a.MaxScore == 0 {
		percentage = 0
	}
	status := []string{}
	if a.Dropped {
		status = append(status, "Dropped")
	}
	if a.Late {
		status = append(status, "Late")
	}
	if a.Missing {
		status = append(status, "Missing")
	}
	if a.Excused {
		status = append(status, "Excused")
	}
	if !a.Graded {
		status = append(status, "Not Graded")
	}
	contrib := a.MaxScore * a.Section.RealWeight / a.Section.TotalPoints
	gradeContrib := a.Score * a.Section.RealWeight / a.Section.TotalPoints
	if a.Section.TotalPoints == 0 {
		contrib = 0
		gradeContrib = 0
	}
	csv.AddRow(a.Name, a.Due.Format("1/2/06 3:04:05 PM"), a.Score, a.MaxScore, percentage, strings.Join(status, ", "), gradeContrib, contrib)
}

type gradedAssignmentSortingMethod int

const (
	gradedAssignmentSortViewOrder    = iota
	gradedAssignmentSortLowestFirst  = iota
	gradedAssignmentSortHighestFirst = iota
)

type gradeSection struct {
	Name        string
	Weight      float64
	TotalScore  float64
	TotalPoints float64
	RealWeight  float64
	Grades      []gradedAssignment
	Order       gradedAssignmentSortingMethod
}

func (s *gradeSection) csv(csv csvgen.CSV) {
	s.Order = gradedAssignmentSortViewOrder
	sort.Sort(s)
	percentage := s.TotalScore * 100 / s.TotalPoints
	if s.TotalPoints == 0 {
		percentage = 0
	}
	sect := csv.AddSection([]interface{}{s.Name}, "", "", s.TotalScore, s.TotalPoints, percentage, "", percentage*s.RealWeight, s.RealWeight*100)
	for _, grade := range s.Grades {
		grade.csv(sect)
	}
}

func (s *gradeSection) countPoints() {
	s.TotalScore = 0
	s.TotalPoints = 0
	for _, grade := range s.Grades {
		if grade.Graded && !grade.Dropped {
			s.TotalScore += grade.Score
			s.TotalPoints += grade.MaxScore
		}
	}
}

func (s *gradeSection) dropGrades(lowest, highest int, protect []int) {
	undroppable := map[int]interface{}{}
	for _, id := range protect {
		undroppable[id] = nil
	}
	if lowest > 0 {
		s.Order = gradedAssignmentSortLowestFirst
		sort.Sort(s)
		for i, grade := range s.Grades {
			if _, ok := undroppable[grade.ID]; !ok && !grade.Excused && grade.Graded {
				s.Grades[i].Dropped = true
				lowest--
				if lowest == 0 {
					break
				}
			}
		}
	}
	if highest > 0 {
		s.Order = gradedAssignmentSortHighestFirst
		sort.Sort(s)
		for i, grade := range s.Grades {
			if _, ok := undroppable[grade.ID]; !ok && !grade.Excused && grade.Graded {
				s.Grades[i].Dropped = true
				highest--
				if highest == 0 {
					break
				}
			}
		}
	}
}

func (s *gradeSection) Len() int {
	return len(s.Grades)
}

func (s *gradeSection) Less(i, j int) bool {
	a := s.Grades[i]
	b := s.Grades[j]
	x := a.Score / a.MaxScore
	y := b.Score / b.MaxScore
	switch s.Order {
	case gradedAssignmentSortViewOrder:
		return a.Position < b.Position
	case gradedAssignmentSortLowestFirst:
		if x == y {
			return a.MaxScore > b.MaxScore
		}
		return x < y
	case gradedAssignmentSortHighestFirst:
		if x == y {
			return a.MaxScore < b.MaxScore
		}
		return x > y
	default:
		panic("Unknown enum constant")
	}
}

func (s *gradeSection) Swap(i, j int) {
	tmp := s.Grades[i]
	s.Grades[i] = s.Grades[j]
	s.Grades[j] = tmp
}

type grades struct {
	Sections []*gradeSection
}

func (g *grades) calculateRealWeights() {
	var totalWeight float64 = 0
	var totalPoints float64 = 0
	for _, s := range g.Sections {
		s.countPoints()
		if s.TotalPoints > 0 {
			totalWeight += s.Weight
			totalPoints += s.TotalPoints
		}
	}
	if totalWeight == 0 {
		for _, s := range g.Sections {
			s.RealWeight = s.TotalPoints / totalPoints
		}
	} else {
		for _, s := range g.Sections {
			s.RealWeight = s.Weight / totalWeight
		}
	}
}

func (g grades) csv(csv csvgen.CSV) {
	var grade float64 = 0
	var score float64 = 0
	var max float64 = 0
	for _, s := range g.Sections {
		s.csv(csv)
		contrib := s.TotalScore * s.RealWeight / s.TotalPoints
		if s.TotalPoints == 0 {
			contrib = 0
		}
		grade += contrib
		score += s.TotalScore
		max += s.TotalPoints
	}
	csv.AddRow("Total", "", "", score, max, 100*grade, "", 100*grade, float64(100))
}

func init() {
	registerCSV("Grades", func(t *task.Task, c *canvas.Canvas, courseId int, csv csvgen.CSV) {
		groups, err := c.AssignmentGroupsListAssignmentGroups(t.CreateProgress(1), []canvas.AssignmentGroupsListAssignmentGroupsInclude{
			canvas.AssignmentGroupsListAssignmentGroupsIncludeAssignments,
			canvas.AssignmentGroupsListAssignmentGroupsIncludeSubmission,
		}, nil, nil, nil, nil, fmt.Sprint(courseId))
		if err != nil {
			panic(err)
		}
		g := &grades{
			Sections: make([]*gradeSection, len(groups)),
		}
		for i, group := range groups {
			grades := make([]gradedAssignment, len(group.Assignments))
			section := &gradeSection{
				Name:   group.Name,
				Weight: group.GroupWeight,
				Grades: grades,
			}
			g.Sections[i] = section
			for i, assignment := range group.Assignments {
				grades[i] = gradedAssignment{
					Name:     assignment.Name,
					Due:      assignment.DueAt,
					ID:       assignment.ID,
					Position: assignment.Position,
					Score:    -1,
					MaxScore: assignment.PointsPossible,
					Dropped:  false,
					Late:     false,
					Missing:  false,
					Excused:  false,
					Graded:   assignment.Submission != nil &&
							  !assignment.Submission.PostedAt.IsZero() &&
							  !assignment.Submission.GradedAt.IsZero() &&
							  *assignment.Submission.WorkflowState != canvas.SubmissionWorkflowStatePendingReview,
					Section:  section,
				}
				if assignment.Submission != nil {
					grades[i].Score = assignment.Submission.Score
					grades[i].Late = assignment.Submission.Late
					grades[i].Missing = assignment.Submission.Missing
					grades[i].Excused = assignment.Submission.Excused
				}
			}
			section.dropGrades(group.Rules.DropLowest, group.Rules.DropHighest, group.Rules.NeverDrop)
		}
		g.calculateRealWeights()
		g.csv(csv)
	}, "Assignment Group", "%s", "Assignment Name", "%s", "Due Date", "%s", "Score", "%.0f", "Max Score", "%.0f",
		"Percentage", "%.2f%%", "Status", "%s", "Total Grade Contribution", "%.2f%%", "Max Grade Contribution", "%.2f%%")
}
