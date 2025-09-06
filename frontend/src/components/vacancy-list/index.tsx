import { useState } from 'react';
import {
  Box,
  Card,
  CardContent,
  Typography,
  Button,
  Grid,
  Chip,
  IconButton,
  Dialog,
  //   CircularProgress,
} from '@mui/material';
import { Edit as EditIcon, Delete as DeleteIcon, Add as AddIcon } from '@mui/icons-material';
// import { useGetVacancies } from '../../services/queries/vacancies';
import type { Vacancy } from '../../types';
import { useDeleteVacancy } from '../../services/mutations/vacancies';
import VacancyForm from '../forms/vacancy-form';

const vacancies: Vacancy[] = [
  {
    id: '1',
    status: 'active',
    name: 'Разработчик Go',
    region: 'Московская область',
    city: 'Москва',
    address: 'ул. Тверская, д. 1',
    work_type: 'полная занятость',
    employment_type: 'штатный сотрудник',
    work_schedule: '5/2',
    income: null,
    salary_max: 250000,
    salary_min: 150000,
    pronounces: 'Приветствуется опыт работы с высоконагруженными системами',
    responsibilities: 'Разработка и поддержка микросервисов на Go',
    requirements: 'Опыт работы от 3 лет, знание PostgreSQL, Redis',
    education: 'Высшее техническое',
    experience: 'от 3 лет',
    special_programs: true,
    computer_skills: true,
    foreign_languages: true,
    language_level: 'B2',
    has_business_trips: false,
    additional_information: 'Возможна частичная удаленная работа',
  },
  {
    id: '2',
    status: 'active',
    name: 'Frontend разработчик (React)',
    region: 'Московская область',
    city: 'Москва',
    address: 'ул. Новый Арбат, д. 15',
    work_type: 'полная занятость',
    employment_type: 'штатный сотрудник',
    work_schedule: '5/2',
    income: null,
    salary_max: 300000,
    salary_min: 180000,
    pronounces: 'Разработка современных веб-приложений',
    responsibilities:
      'Разработка пользовательского интерфейса на React, оптимизация производительности',
    requirements: 'Опыт работы с React от 2 лет, TypeScript, Redux, Webpack',
    education: 'Высшее техническое',
    experience: 'от 2 лет',
    special_programs: true,
    computer_skills: true,
    foreign_languages: true,
    language_level: 'B1',
    has_business_trips: true,
    additional_information: 'Гибкий график, удаленная работа',
  },
  {
    id: '3',
    status: 'draft',
    name: 'Backend разработчик (Python)',
    region: 'Санкт-Петербург',
    city: 'Санкт-Петербург',
    address: 'Невский проспект, д. 100',
    work_type: 'полная занятость',
    employment_type: 'штатный сотрудник',
    work_schedule: '5/2',
    income: null,
    salary_max: 280000,
    salary_min: 160000,
    pronounces: 'Разработка высоконагруженных систем',
    responsibilities: 'Разработка API, работа с базами данных, оптимизация запросов',
    requirements: 'Python 3.8+, Django/FastAPI, PostgreSQL, Docker',
    education: 'Высшее техническое',
    experience: 'от 3 лет',
    special_programs: true,
    computer_skills: true,
    foreign_languages: false,
    language_level: '',
    has_business_trips: false,
    additional_information: 'Офис в центре города',
  },
  {
    id: '4',
    status: 'active',
    name: 'DevOps инженер',
    region: 'Новосибирская область',
    city: 'Новосибирск',
    address: 'ул. Ленина, д. 25',
    work_type: 'полная занятость',
    employment_type: 'штатный сотрудник',
    work_schedule: '5/2',
    income: null,
    salary_max: 350000,
    salary_min: 220000,
    pronounces: 'Построение CI/CD процессов',
    responsibilities: 'Настройка и поддержка инфраструктуры, автоматизация деплоя',
    requirements: 'Kubernetes, Docker, AWS/GCP, Terraform, Ansible',
    education: 'Высшее техническое',
    experience: 'от 4 лет',
    special_programs: true,
    computer_skills: true,
    foreign_languages: true,
    language_level: 'B2',
    has_business_trips: true,
    additional_information: 'Возможны командировки',
  },
  {
    id: '5',
    status: 'closed',
    name: 'Data Scientist',
    region: 'Московская область',
    city: 'Москва',
    address: 'ул. Пушкина, д. 10',
    work_type: 'полная занятость',
    employment_type: 'штатный сотрудник',
    work_schedule: '5/2',
    income: null,
    salary_max: 320000,
    salary_min: 200000,
    pronounces: 'Работа с большими данными',
    responsibilities: 'Разработка ML моделей, анализ данных, построение предсказательных моделей',
    requirements: 'Python, Pandas, NumPy, Scikit-learn, SQL',
    education: 'Высшее техническое или математическое',
    experience: 'от 3 лет',
    special_programs: true,
    computer_skills: true,
    foreign_languages: true,
    language_level: 'B2',
    has_business_trips: false,
    additional_information: 'Исследовательский подход к работе',
  },
  {
    id: '6',
    status: 'active',
    name: 'UX/UI дизайнер',
    region: 'Московская область',
    city: 'Москва',
    address: 'ул. Садовая, д. 5',
    work_type: 'удаленная работа',
    employment_type: 'штатный сотрудник',
    work_schedule: 'гибкий график',
    income: null,
    salary_max: 220000,
    salary_min: 120000,
    pronounces: 'Создание пользовательских интерфейсов',
    responsibilities: 'Проектирование интерфейсов, создание прототипов, работа с дизайн-системами',
    requirements: 'Figma, Adobe XD, знание принципов UX, портфолио',
    education: 'Высшее дизайнерское',
    experience: 'от 2 лет',
    special_programs: true,
    computer_skills: true,
    foreign_languages: true,
    language_level: 'B1',
    has_business_trips: false,
    additional_information: 'Полностью удаленная работа',
  },
  {
    id: '7',
    status: 'active',
    name: 'Менеджер проектов',
    region: 'Московская область',
    city: 'Москва',
    address: 'ул. Тверская, д. 20',
    work_type: 'полная занятость',
    employment_type: 'штатный сотрудник',
    work_schedule: '5/2',
    income: null,
    salary_max: 280000,
    salary_min: 180000,
    pronounces: 'Управление IT проектами',
    responsibilities: 'Планирование проектов, управление командой, контроль сроков и бюджета',
    requirements: 'Опыт управления проектами от 3 лет, знание Agile/Scrum',
    education: 'Высшее',
    experience: 'от 3 лет',
    special_programs: true,
    computer_skills: true,
    foreign_languages: true,
    language_level: 'B2',
    has_business_trips: true,
    additional_information: 'Офис в центре, соцпакет',
  },
  {
    id: '8',
    status: 'active',
    name: 'Тестировщик (QA)',
    region: 'Екатеринбург',
    city: 'Екатеринбург',
    address: 'ул. Малышева, д. 30',
    work_type: 'полная занятость',
    employment_type: 'штатный сотрудник',
    work_schedule: '5/2',
    income: null,
    salary_max: 180000,
    salary_min: 100000,
    pronounces: 'Тестирование веб-приложений',
    responsibilities: 'Написание тест-кейсов, ручное и автоматизированное тестирование',
    requirements: 'Опыт тестирования от 1 года, понимание жизненного цикла ПО',
    education: 'Высшее техническое',
    experience: 'от 1 года',
    special_programs: true,
    computer_skills: true,
    foreign_languages: false,
    language_level: '',
    has_business_trips: false,
    additional_information: 'Обучение за счет компании',
  },
];

const VacancyList = () => {
  //   const { data:vacancies, isLoading, isError, error } = useGetVacancies();
  const { mutate: deleteVacancy } = useDeleteVacancy();

  const [isFormOpen, setIsFormOpen] = useState(false);
  const [editingVacancy, setEditingVacancy] = useState<Vacancy | null>(null);

  const handleAdd = () => {
    setEditingVacancy(null);
    setIsFormOpen(true);
  };

  const handleEdit = (vacancy: Vacancy) => {
    setEditingVacancy(vacancy);
    setIsFormOpen(true);
  };

  const handleDelete = (id: string) => {
    if (window.confirm('Вы уверены, что хотите удалить эту вакансию?')) {
      deleteVacancy(id);
    }
  };

  const handleCloseForm = () => {
    setIsFormOpen(false);
    setEditingVacancy(null);
  };

  //   if (isLoading) {
  // return (
  //   <Box display="flex" justifyContent="center" alignItems="center" minHeight="200px">
  // <CircularProgress />
  //   </Box>
  // );
  //   }

  //   if(isError){
  // return (
  // <Box>
  // <Typography>{`Произошла ошибка: ${error}`}</Typography>
  // </Box>
  // );
  //   }

  return (
    <Box>
      <Box display='flex' justifyContent='space-between' alignItems='center' mb={3}>
        <Typography variant='h4'>Список вакансий</Typography>
        <Button variant='contained' startIcon={<AddIcon />} onClick={handleAdd}>
          Добавить вакансию
        </Button>
      </Box>

      <Grid container spacing={3}>
        {vacancies?.map((vacancy) => (
          <Grid sx={{ xs: 12, md: 6, lg: 4 }} key={vacancy.id}>
            <Card sx={{ height: '100%' }}>
              <CardContent>
                <Box display='flex' justifyContent='space-between' alignItems='flex-start' mb={2}>
                  <Typography variant='h6' gutterBottom>
                    {vacancy.name}
                  </Typography>
                  <Chip
                    label={vacancy.status}
                    color={vacancy.status === 'active' ? 'success' : 'default'}
                    size='small'
                  />
                </Box>

                <Typography variant='body2' color='text.secondary' gutterBottom>
                  {vacancy.city}, {vacancy.region}
                </Typography>

                <Typography variant='body2' gutterBottom>
                  Зарплата: {vacancy.salary_min} - {vacancy.salary_max} руб.
                </Typography>

                <Typography variant='body2' color='text.secondary' paragraph>
                  {vacancy.work_type} • {vacancy.employment_type}
                </Typography>

                <Box display='flex' justifyContent='space-between' alignItems='center'>
                  <Typography variant='caption' color='text.secondary'>
                    Опыт: {vacancy.experience}
                  </Typography>
                  <Box>
                    <IconButton
                      size='small'
                      onClick={() => handleEdit(vacancy)}
                      //   disabled={isDeleting}
                    >
                      <EditIcon />
                    </IconButton>
                    <IconButton
                      size='small'
                      onClick={() => handleDelete(vacancy.id!)}
                      //   disabled={isDeleting}
                      color='error'
                    >
                      <DeleteIcon />
                    </IconButton>
                  </Box>
                </Box>
              </CardContent>
            </Card>
          </Grid>
        ))}
      </Grid>

      <Dialog open={isFormOpen} onClose={handleCloseForm} maxWidth='md' fullWidth>
        <VacancyForm
          vacancy={editingVacancy}
          onClose={handleCloseForm}
          onSuccess={handleCloseForm}
        />
      </Dialog>
    </Box>
  );
};

export default VacancyList;
