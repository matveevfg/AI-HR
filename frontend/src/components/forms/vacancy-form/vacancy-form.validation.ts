import { z } from 'zod';

export const vacancySchema = z.object({
  status: z.string().min(1, 'Статус обязателен'),
  name: z.string().min(1, 'Название вакансии обязательно'),
  region: z.string().min(1, 'Регион обязателен'),
  city: z.string().min(1, 'Город обязателен'),
  address: z.string().min(1, 'Адрес обязателен'),
  work_type: z.string().min(1, 'Тип работы обязателен'),
  employment_type: z.string().min(1, 'Тип занятости обязателен'),
  work_schedule: z.string().min(1, 'График работы обязателен'),
  income: z.string().nullable(),
  salary_max: z.number().min(0, 'Максимальная зарплата должна быть положительной'),
  salary_min: z.number().min(0, 'Минимальная зарплата должна быть положительной'),
  pronounces: z.string(),
  responsibilities: z.string().min(1, 'Обязанности обязательны'),
  requirements: z.string().min(1, 'Требования обязательны'),
  education: z.string(),
  experience: z.string(),
  special_programs: z.boolean(),
  computer_skills: z.boolean(),
  foreign_languages: z.boolean(),
  language_level: z.string(),
  has_business_trips: z.boolean(),
  additional_information: z.string(),
});

export type VacancyFormData = z.infer<typeof vacancySchema>;
