import { useQuery } from '@tanstack/react-query';
import { getVacancies } from '../../api/vacancies';

export const useGetVacancies = () =>
  useQuery({
    queryKey: ['get-vacancies'],
    queryFn: getVacancies,
  });
