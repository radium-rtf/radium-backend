// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package db

import (
	"context"
	"database/sql"

	"gorm.io/gorm"

	"gorm.io/gen"

	"gorm.io/plugin/dbresolver"
)

func Use(db *gorm.DB, opts ...gen.DOOption) *Query {
	return &Query{
		db:                       db,
		Answer:                   newAnswer(db, opts...),
		AnswerSection:            newAnswerSection(db, opts...),
		AnswerSectionAnswer:      newAnswerSectionAnswer(db, opts...),
		ChoiceSection:            newChoiceSection(db, opts...),
		ChoiceSectionAnswer:      newChoiceSectionAnswer(db, opts...),
		CodeSection:              newCodeSection(db, opts...),
		CodeSectionAnswer:        newCodeSectionAnswer(db, opts...),
		Course:                   newCourse(db, opts...),
		Group:                    newGroup(db, opts...),
		Link:                     newLink(db, opts...),
		Module:                   newModule(db, opts...),
		MultiChoiceSection:       newMultiChoiceSection(db, opts...),
		MultichoiceSectionAnswer: newMultichoiceSectionAnswer(db, opts...),
		Page:                     newPage(db, opts...),
		Review:                   newReview(db, opts...),
		Section:                  newSection(db, opts...),
		Session:                  newSession(db, opts...),
		ShortAnswerSection:       newShortAnswerSection(db, opts...),
		ShortAnswerSectionAnswer: newShortAnswerSectionAnswer(db, opts...),
		Teacher:                  newTeacher(db, opts...),
		TeacherCourse:            newTeacherCourse(db, opts...),
		TextSection:              newTextSection(db, opts...),
		User:                     newUser(db, opts...),
	}
}

type Query struct {
	db *gorm.DB

	Answer                   answer
	AnswerSection            answerSection
	AnswerSectionAnswer      answerSectionAnswer
	ChoiceSection            choiceSection
	ChoiceSectionAnswer      choiceSectionAnswer
	CodeSection              codeSection
	CodeSectionAnswer        codeSectionAnswer
	Course                   course
	Group                    group
	Link                     link
	Module                   module
	MultiChoiceSection       multiChoiceSection
	MultichoiceSectionAnswer multichoiceSectionAnswer
	Page                     page
	Review                   review
	Section                  section
	Session                  session
	ShortAnswerSection       shortAnswerSection
	ShortAnswerSectionAnswer shortAnswerSectionAnswer
	Teacher                  teacher
	TeacherCourse            teacherCourse
	TextSection              textSection
	User                     user
}

func (q *Query) Available() bool { return q.db != nil }

func (q *Query) clone(db *gorm.DB) *Query {
	return &Query{
		db:                       db,
		Answer:                   q.Answer.clone(db),
		AnswerSection:            q.AnswerSection.clone(db),
		AnswerSectionAnswer:      q.AnswerSectionAnswer.clone(db),
		ChoiceSection:            q.ChoiceSection.clone(db),
		ChoiceSectionAnswer:      q.ChoiceSectionAnswer.clone(db),
		CodeSection:              q.CodeSection.clone(db),
		CodeSectionAnswer:        q.CodeSectionAnswer.clone(db),
		Course:                   q.Course.clone(db),
		Group:                    q.Group.clone(db),
		Link:                     q.Link.clone(db),
		Module:                   q.Module.clone(db),
		MultiChoiceSection:       q.MultiChoiceSection.clone(db),
		MultichoiceSectionAnswer: q.MultichoiceSectionAnswer.clone(db),
		Page:                     q.Page.clone(db),
		Review:                   q.Review.clone(db),
		Section:                  q.Section.clone(db),
		Session:                  q.Session.clone(db),
		ShortAnswerSection:       q.ShortAnswerSection.clone(db),
		ShortAnswerSectionAnswer: q.ShortAnswerSectionAnswer.clone(db),
		Teacher:                  q.Teacher.clone(db),
		TeacherCourse:            q.TeacherCourse.clone(db),
		TextSection:              q.TextSection.clone(db),
		User:                     q.User.clone(db),
	}
}

func (q *Query) ReadDB() *Query {
	return q.ReplaceDB(q.db.Clauses(dbresolver.Read))
}

func (q *Query) WriteDB() *Query {
	return q.ReplaceDB(q.db.Clauses(dbresolver.Write))
}

func (q *Query) ReplaceDB(db *gorm.DB) *Query {
	return &Query{
		db:                       db,
		Answer:                   q.Answer.replaceDB(db),
		AnswerSection:            q.AnswerSection.replaceDB(db),
		AnswerSectionAnswer:      q.AnswerSectionAnswer.replaceDB(db),
		ChoiceSection:            q.ChoiceSection.replaceDB(db),
		ChoiceSectionAnswer:      q.ChoiceSectionAnswer.replaceDB(db),
		CodeSection:              q.CodeSection.replaceDB(db),
		CodeSectionAnswer:        q.CodeSectionAnswer.replaceDB(db),
		Course:                   q.Course.replaceDB(db),
		Group:                    q.Group.replaceDB(db),
		Link:                     q.Link.replaceDB(db),
		Module:                   q.Module.replaceDB(db),
		MultiChoiceSection:       q.MultiChoiceSection.replaceDB(db),
		MultichoiceSectionAnswer: q.MultichoiceSectionAnswer.replaceDB(db),
		Page:                     q.Page.replaceDB(db),
		Review:                   q.Review.replaceDB(db),
		Section:                  q.Section.replaceDB(db),
		Session:                  q.Session.replaceDB(db),
		ShortAnswerSection:       q.ShortAnswerSection.replaceDB(db),
		ShortAnswerSectionAnswer: q.ShortAnswerSectionAnswer.replaceDB(db),
		Teacher:                  q.Teacher.replaceDB(db),
		TeacherCourse:            q.TeacherCourse.replaceDB(db),
		TextSection:              q.TextSection.replaceDB(db),
		User:                     q.User.replaceDB(db),
	}
}

type queryCtx struct {
	Answer                   IAnswerDo
	AnswerSection            IAnswerSectionDo
	AnswerSectionAnswer      IAnswerSectionAnswerDo
	ChoiceSection            IChoiceSectionDo
	ChoiceSectionAnswer      IChoiceSectionAnswerDo
	CodeSection              ICodeSectionDo
	CodeSectionAnswer        ICodeSectionAnswerDo
	Course                   ICourseDo
	Group                    IGroupDo
	Link                     ILinkDo
	Module                   IModuleDo
	MultiChoiceSection       IMultiChoiceSectionDo
	MultichoiceSectionAnswer IMultichoiceSectionAnswerDo
	Page                     IPageDo
	Review                   IReviewDo
	Section                  ISectionDo
	Session                  ISessionDo
	ShortAnswerSection       IShortAnswerSectionDo
	ShortAnswerSectionAnswer IShortAnswerSectionAnswerDo
	Teacher                  ITeacherDo
	TeacherCourse            ITeacherCourseDo
	TextSection              ITextSectionDo
	User                     IUserDo
}

func (q *Query) WithContext(ctx context.Context) *queryCtx {
	return &queryCtx{
		Answer:                   q.Answer.WithContext(ctx),
		AnswerSection:            q.AnswerSection.WithContext(ctx),
		AnswerSectionAnswer:      q.AnswerSectionAnswer.WithContext(ctx),
		ChoiceSection:            q.ChoiceSection.WithContext(ctx),
		ChoiceSectionAnswer:      q.ChoiceSectionAnswer.WithContext(ctx),
		CodeSection:              q.CodeSection.WithContext(ctx),
		CodeSectionAnswer:        q.CodeSectionAnswer.WithContext(ctx),
		Course:                   q.Course.WithContext(ctx),
		Group:                    q.Group.WithContext(ctx),
		Link:                     q.Link.WithContext(ctx),
		Module:                   q.Module.WithContext(ctx),
		MultiChoiceSection:       q.MultiChoiceSection.WithContext(ctx),
		MultichoiceSectionAnswer: q.MultichoiceSectionAnswer.WithContext(ctx),
		Page:                     q.Page.WithContext(ctx),
		Review:                   q.Review.WithContext(ctx),
		Section:                  q.Section.WithContext(ctx),
		Session:                  q.Session.WithContext(ctx),
		ShortAnswerSection:       q.ShortAnswerSection.WithContext(ctx),
		ShortAnswerSectionAnswer: q.ShortAnswerSectionAnswer.WithContext(ctx),
		Teacher:                  q.Teacher.WithContext(ctx),
		TeacherCourse:            q.TeacherCourse.WithContext(ctx),
		TextSection:              q.TextSection.WithContext(ctx),
		User:                     q.User.WithContext(ctx),
	}
}

func (q *Query) Transaction(fc func(tx *Query) error, opts ...*sql.TxOptions) error {
	return q.db.Transaction(func(tx *gorm.DB) error { return fc(q.clone(tx)) }, opts...)
}

func (q *Query) Begin(opts ...*sql.TxOptions) *QueryTx {
	tx := q.db.Begin(opts...)
	return &QueryTx{Query: q.clone(tx), Error: tx.Error}
}

type QueryTx struct {
	*Query
	Error error
}

func (q *QueryTx) Commit() error {
	return q.db.Commit().Error
}

func (q *QueryTx) Rollback() error {
	return q.db.Rollback().Error
}

func (q *QueryTx) SavePoint(name string) error {
	return q.db.SavePoint(name).Error
}

func (q *QueryTx) RollbackTo(name string) error {
	return q.db.RollbackTo(name).Error
}
