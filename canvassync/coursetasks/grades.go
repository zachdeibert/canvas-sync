package coursetasks

import (
	"fmt"
	"io/ioutil"
	"path"
	"sort"
	"strings"
	"time"

	"github.com/zachdeibert/canvas-sync/canvas"
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
}

type gradedAssignmentSortingMethod int

const (
	gradedAssignmentSortViewOrder    = iota
	gradedAssignmentSortLowestFirst  = iota
	gradedAssignmentSortHighestFirst = iota
)

type gradedAssignmentSorter struct {
	grades []gradedAssignment
	mode   gradedAssignmentSortingMethod
}

func (s *gradedAssignmentSorter) Len() int {
	return len(s.grades)
}

func (s *gradedAssignmentSorter) Less(i, j int) bool {
	a := s.grades[i]
	b := s.grades[j]
	x := a.Score / a.MaxScore
	y := b.Score / b.MaxScore
	switch s.mode {
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

func (s *gradedAssignmentSorter) Swap(i, j int) {
	tmp := s.grades[i]
	s.grades[i] = s.grades[j]
	s.grades[j] = tmp
}

func init() {
	register("Grades", func(t *task.Task, c *canvas.Canvas, db string, courseId int, finish func()) {
		groups, err := c.AssignmentGroupsListAssignmentGroups(t.CreateProgress(1), []canvas.AssignmentGroupsListAssignmentGroupsInclude{
			canvas.AssignmentGroupsListAssignmentGroupsIncludeAssignments,
			canvas.AssignmentGroupsListAssignmentGroupsIncludeSubmission,
		}, nil, nil, nil, nil, fmt.Sprint(courseId))
		if err != nil {
			if e, ok := err.(canvas.InvalidStatusCodeError); ok && e.Code == 401 {
				finish()
			} else {
				panic(err)
			}
		}
		str := &strings.Builder{}
		fmt.Fprintln(str, "Assignment Group,Assignment Name,Due Date,Score,Max Score,Percentage,Status,Total Grade Contribution,Max Grade Contribution")
		var totalWeight float64 = 0
		var missingWeight float64 = 0
		for _, group := range groups {
			totalWeight += float64(group.GroupWeight)
		}
		var totalGrade float64 = 0
		var courseTotalPoints float64 = 0
		gradeList := make([][]gradedAssignment, len(groups))
		for idx, group := range groups {
			undroppable := map[int]interface{}{}
			for _, id := range group.Rules.NeverDrop {
				undroppable[id] = nil
			}
			grades := make([]gradedAssignment, len(group.Assignments))
			gradeList[idx] = grades
			for i, assignment := range group.Assignments {
				graded := !assignment.Submission.PostedAt.IsZero() &&
					!assignment.Submission.GradedAt.IsZero() &&
					*assignment.Submission.WorkflowState != canvas.SubmissionWorkflowStatePendingReview
				grades[i] = gradedAssignment{
					Name:     assignment.Name,
					Due:      assignment.DueAt,
					ID:       assignment.ID,
					Position: assignment.Position,
					Score:    assignment.Submission.Score,
					MaxScore: assignment.PointsPossible,
					Dropped:  false,
					Late:     assignment.Submission.Late,
					Missing:  assignment.Submission.Missing,
					Excused:  assignment.Submission.Excused,
					Graded:   graded,
				}
			}
			sorter := &gradedAssignmentSorter{
				grades: grades,
			}
			if group.Rules.DropLowest > 0 {
				sorter.mode = gradedAssignmentSortLowestFirst
				sort.Sort(sorter)
				count := group.Rules.DropLowest
				for i, grade := range grades {
					if _, ok := undroppable[grade.ID]; !ok && !grades[i].Excused && grades[i].Graded {
						grades[i].Dropped = true
						count--
						if count == 0 {
							break
						}
					}
				}
			}
			if group.Rules.DropHighest > 0 {
				sorter.mode = gradedAssignmentSortHighestFirst
				sort.Sort(sorter)
				count := group.Rules.DropHighest
				for i, grade := range grades {
					if _, ok := undroppable[grade.ID]; !ok && !grades[i].Excused && grades[i].Graded {
						grades[i].Dropped = true
						count--
						if count == 0 {
							break
						}
					}
				}
			}
			sorter.mode = gradedAssignmentSortViewOrder
			sort.Sort(sorter)
			for _, grade := range grades {
				if grade.Graded && !grade.Dropped {
					courseTotalPoints += grade.MaxScore
				}
			}
		}
		for idx, group := range groups {
			grades := gradeList[idx]
			var totalScore float64 = 0
			var totalMaxScore float64 = 0
			for _, grade := range grades {
				if grade.Graded && !grade.Dropped {
					totalScore += grade.Score
					totalMaxScore += grade.MaxScore
				}
			}
			contribution := float64(group.GroupWeight) / totalWeight * 100
			if totalWeight == 0 {
				contribution = 100 * totalMaxScore / courseTotalPoints
			}
			groupTotal := totalScore / totalMaxScore
			if totalMaxScore == 0 {
				groupTotal = 0
				if contribution != 0 {
					missingWeight += contribution / 100
				}
			}
			totalGrade += contribution * groupTotal
			fmt.Fprintf(str, "\"%s\",,,%.0f,%.0f,%.2f%%,,%.2f%%,%.2f%%\n", group.Name, totalScore, totalMaxScore, groupTotal*100, contribution*groupTotal, contribution)
			if totalMaxScore == 0 && totalScore == 0 {
				totalMaxScore = 1
			}
			for _, grade := range grades {
				frac := grade.Score / grade.MaxScore
				if grade.MaxScore == 0 {
					frac = 0
				}
				gradeContribution := grade.MaxScore / totalMaxScore * contribution
				status := []string{}
				if grade.Dropped {
					status = append(status, "Dropped")
				}
				if grade.Late {
					status = append(status, "Late")
				}
				if grade.Missing {
					status = append(status, "Missing")
				}
				if grade.Excused {
					status = append(status, "Excused")
				}
				if !grade.Graded {
					status = append(status, "Not Graded")
				}
				fmt.Fprintf(str, ",\"%s\",%s,%.0f,%.0f,%.2f%%,\"%s\",%.2f%%,%.2f%%\n", grade.Name, grade.Due.Format(""), grade.Score, grade.MaxScore, frac*100, strings.Join(status, ", "), grade.Score/totalMaxScore*contribution, gradeContribution)
			}
		}
		fmt.Fprintf(str, "Total,,,,,%.2f%%,,,\n", totalGrade/(1-missingWeight))
		if err := ioutil.WriteFile(path.Join(db, "Grades.csv"), []byte(str.String()), 0644); err != nil {
			panic(err)
		}
		finish()
	})
}
