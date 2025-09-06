import api from '..';
import type { Vacancy } from '../../../types';

export const getVacancies = async () => {
  const response = await api.get('/vacancies');
  return response.data;
};

export const createVacancy = async (data: Vacancy) => {
  const response = await api.post('/vacancies', data);
  return response.data;
};

export const updateVacancy = async (data: Vacancy) => {
  const response = await api.put(`/vacancies/${data.id}`, data);
  return response.data;
};

export const deleteVacancy = async (id: string) => {
  await api.delete(`/vacancies/${id}`);
};
