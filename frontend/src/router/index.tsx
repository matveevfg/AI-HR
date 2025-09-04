import { createBrowserRouter, Navigate } from 'react-router-dom';
import MainPage from '../pages/main';
import VacanciesPage from '../pages/vacancies';
import CandidatesPage from '../pages/candidates';
import InterviewsPage from '../pages/interviews';

const router = createBrowserRouter([
  {
    path: '*',
    element: <Navigate to='/' />,
  },
  {
    path: '/',
    element: <MainPage />,
  },
  {
    path: '/vacancies',
    element: <VacanciesPage />,
  },
  {
    path: '/candidates',
    element: <CandidatesPage />,
  },
  {
    path: '/interviews',
    element: <InterviewsPage />,
  },
]);

export default router;
