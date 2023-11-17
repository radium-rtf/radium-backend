package usecase

import (
	"context"
	"errors"
	"github.com/google/uuid"
	"github.com/radium-rtf/radium-backend/internal/usecase/repo/postgres"
	"github.com/radium-rtf/radium-backend/internal/usecase/repo/repoerr"

	"github.com/radium-rtf/radium-backend/internal/entity"
)

type SectionUseCase struct {
	section postgres.Section
	page    postgres.Page
}

func NewSectionUseCase(sectionRepo postgres.Section, page postgres.Page) SectionUseCase {
	return SectionUseCase{section: sectionRepo, page: page}
}

func (uc SectionUseCase) Create(ctx context.Context, section *entity.Section, editorId uuid.UUID) (*entity.Section, error) {
	course, err := uc.page.GetCourseByPageId(ctx, section.PageId)
	if err != nil {
		return nil, err
	}

	if !course.CanEdit(editorId) {
		return nil, cantEditCourse
	}

	last, err := uc.section.GetLastSection(ctx, section.PageId)
	if err != nil && !errors.Is(err, repoerr.SectionNotFound) {
		return nil, err
	}

	section.Order = 1
	if !errors.Is(err, repoerr.SectionNotFound) {
		section.Order = last.Order + 1
	}

	return uc.section.Create(ctx, section)
}

func (uc SectionUseCase) Delete(ctx context.Context, id, editorId uuid.UUID, isSoft bool) error {
	canEditErr := uc.canEdit(ctx, id, editorId)
	if canEditErr != nil {
		return canEditErr
	}
	return uc.section.Delete(ctx, id, isSoft)
}

func (uc SectionUseCase) Update(ctx context.Context, section *entity.Section, editorId uuid.UUID) (*entity.Section, error) {
	canEditErr := uc.canEdit(ctx, section.Id, editorId)
	if canEditErr != nil {
		return nil, canEditErr
	}
	return uc.section.Update(ctx, section)
}

func (uc SectionUseCase) canEdit(ctx context.Context, id, editorId uuid.UUID) error {
	course, err := uc.section.GetCourseBySectionId(ctx, id)
	if err != nil {
		return err
	}

	if !course.CanEdit(editorId) {
		return cantEditCourse
	}
	return nil
}
