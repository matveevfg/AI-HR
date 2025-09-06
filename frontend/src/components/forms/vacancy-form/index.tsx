import {
  DialogTitle,
  DialogContent,
  DialogActions,
  Button,
  TextField,
  FormControlLabel,
  Checkbox,
  Grid,
  Box,
} from '@mui/material';
import { useForm } from 'react-hook-form';
import { zodResolver } from '@hookform/resolvers/zod';
import type { Vacancy, VacancyFormData } from '../../../types';
import { vacancySchema } from './vacancy-form.validation';
import { useCreateVacancy, useUpdateVacancy } from '../../../services/mutations/vacancies';

interface VacancyFormProps {
  vacancy?: Vacancy | null;
  onClose: () => void;
  onSuccess: () => void;
}

const VacancyForm = ({ vacancy, onClose, onSuccess }: VacancyFormProps) => {
  const { mutate: updateVacancy, isSuccess: isUpdating } = useUpdateVacancy();
  const { mutate: createVacancy, isSuccess: isCreating } = useCreateVacancy();

  const {
    register,
    handleSubmit,
    formState: { errors },
    watch,
  } = useForm<VacancyFormData>({
    resolver: zodResolver(vacancySchema),
    defaultValues: vacancy || {
      status: 'active',
      special_programs: false,
      computer_skills: false,
      foreign_languages: false,
      has_business_trips: false,
    },
  });

  const foreignLanguages = watch('foreign_languages');

  const onSubmit = (data: VacancyFormData) => {
    if (vacancy?.id) {
      console.log(data);
      updateVacancy({ id: vacancy.id, data });
    } else {
      console.log(data);
      createVacancy(data);
    }
    onSuccess();
  };

  const isSubmitting = isCreating || isUpdating;

  return (
    <form onSubmit={handleSubmit(onSubmit)}>
      <DialogTitle>{vacancy ? 'Редактировать вакансию' : 'Создать вакансию'}</DialogTitle>

      <DialogContent>
        <Box sx={{ pt: 2 }}>
          <Grid container spacing={2}>
            <Grid sx={{ xs: 12 }}>
              <TextField
                fullWidth
                label='Название вакансии'
                {...register('name')}
                error={!!errors.name}
                helperText={errors.name?.message}
              />
            </Grid>

            <Grid sx={{ xs: 12, sm: 6 }}>
              <TextField
                fullWidth
                label='Регион'
                {...register('region')}
                error={!!errors.region}
                helperText={errors.region?.message}
              />
            </Grid>

            <Grid sx={{ xs: 12, sm: 6 }}>
              <TextField
                fullWidth
                label='Город'
                {...register('city')}
                error={!!errors.city}
                helperText={errors.city?.message}
              />
            </Grid>

            <Grid sx={{ xs: 12 }}>
              <TextField
                fullWidth
                label='Адрес'
                {...register('address')}
                error={!!errors.address}
                helperText={errors.address?.message}
              />
            </Grid>

            <Grid sx={{ xs: 12, sm: 6 }}>
              <TextField
                fullWidth
                label='Минимальная зарплата'
                type='number'
                {...register('salary_min', { valueAsNumber: true })}
                error={!!errors.salary_min}
                helperText={errors.salary_min?.message}
              />
            </Grid>

            <Grid sx={{ xs: 12, sm: 6 }}>
              <TextField
                fullWidth
                label='Максимальная зарплата'
                type='number'
                {...register('salary_max', { valueAsNumber: true })}
                error={!!errors.salary_max}
                helperText={errors.salary_max?.message}
              />
            </Grid>

            <Grid sx={{ xs: 12 }}>
              <TextField
                fullWidth
                multiline
                rows={3}
                label='Обязанности'
                {...register('responsibilities')}
                error={!!errors.responsibilities}
                helperText={errors.responsibilities?.message}
              />
            </Grid>

            <Grid sx={{ xs: 12 }}>
              <TextField
                fullWidth
                multiline
                rows={3}
                label='Требования'
                {...register('requirements')}
                error={!!errors.requirements}
                helperText={errors.requirements?.message}
              />
            </Grid>

            <Grid sx={{ xs: 12, sm: 6 }}>
              <FormControlLabel
                control={<Checkbox {...register('special_programs')} />}
                label='Специальные программы'
              />
            </Grid>

            <Grid sx={{ xs: 12, sm: 6 }}>
              <FormControlLabel
                control={<Checkbox {...register('computer_skills')} />}
                label='Компьютерные навыки'
              />
            </Grid>

            <Grid sx={{ xs: 12, sm: 6 }}>
              <FormControlLabel
                control={<Checkbox {...register('foreign_languages')} />}
                label='Иностранные языки'
              />
            </Grid>

            {foreignLanguages && (
              <Grid sx={{ xs: 12, sm: 6 }}>
                <TextField
                  fullWidth
                  label='Уровень языка'
                  {...register('language_level')}
                  error={!!errors.language_level}
                  helperText={errors.language_level?.message}
                />
              </Grid>
            )}
          </Grid>
        </Box>
      </DialogContent>

      <DialogActions>
        <Button onClick={onClose} disabled={isSubmitting}>
          Отмена
        </Button>
        <Button type='submit' variant='contained' disabled={isSubmitting}>
          {isSubmitting ? 'Сохранение...' : vacancy ? 'Обновить' : 'Создать'}
        </Button>
      </DialogActions>
    </form>
  );
};

export default VacancyForm;
