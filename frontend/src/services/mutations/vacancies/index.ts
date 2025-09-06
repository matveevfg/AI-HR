import { useMutation } from '@tanstack/react-query';
import { createVacancy, deleteVacancy, updateVacancy } from '../../api/vacancies';
import { queryClient } from '../../../main';

export const useCreateVacancy = () =>
  useMutation({
    mutationFn: createVacancy,
    onSuccess: () => queryClient.invalidateQueries({ queryKey: ['vacancies'] }),
  });

export const useUpdateVacancy = () =>
  useMutation({
    mutationKey: ['update-vacancy'],
    mutationFn: updateVacancy,
    onSuccess: () => queryClient.invalidateQueries({ queryKey: ['vacancies'] }),
  });

export const useDeleteVacancy = () =>
  useMutation({
    mutationFn: deleteVacancy,
    onSuccess: () => queryClient.invalidateQueries({ queryKey: ['vacancies'] }),
  });
