import { createBrowserRouter, Navigate } from 'react-router-dom';
import MainPage from '../pages/main';

const router = createBrowserRouter([
  {
    path: '*',
    element: <Navigate to='/' />,
  },
  {
    path: '/',
    element: <MainPage />,
  },
]);

export default router;
