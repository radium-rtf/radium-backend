create view slide_sections_view as (
  select 
    *, 
    array_to_json(
      (
        select 
          (
            array_agg(
              row_to_json(text.*)
            )
          ) as texts 
        from 
          (
            select 
              public.sections_text.* 
            from 
              sections_text 
            where 
              public.sections_text.slide_id = slides.id
          ) as text
      ) || (
        select 
          array_to_json(
            array_agg(
              row_to_json(question.*)
            )
          ) as questions 
        from 
          (
            select 
              public.sections_question.* 
            from 
              sections_question 
            where 
              public.sections_question.slide_id = slides.id
          ) as question
      ) || (
        select 
          array_to_json(
            array_agg(
              row_to_json(choice.*)
            )
          ) as choices 
        from 
          (
            select 
              public.sections_choice.* 
            from 
              sections_choice 
            where 
              public.sections_choice.slide_id = slides.id
          ) as choice
      ) || (
        select 
          array_to_json(
            array_agg(
              row_to_json(multichoice.*)
            )
          ) as multichoices 
        from 
          (
            select 
              public.sections_multi_choice.* 
            from 
              sections_multi_choice 
            where 
              public.sections_multi_choice.slide_id = slides.id
          ) as multichoice
      )
    ) as sections
  from 
    slides
);
