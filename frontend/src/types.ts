export interface Vacancy {
  id?: string;
  status: string;
  name: string;
  region: string;
  city: string;
  address: string;
  work_type: string;
  employment_type: string;
  work_schedule: string;
  income: string | null;
  salary_max: number;
  salary_min: number;
  pronounces: string;
  responsibilities: string;
  requirements: string;
  education: string;
  experience: string;
  special_programs: boolean;
  computer_skills: boolean;
  foreign_languages: boolean;
  language_level: string;
  has_business_trips: boolean;
  additional_information: string;
}

export type VacancyFormData = Omit<Vacancy, 'id'>;
