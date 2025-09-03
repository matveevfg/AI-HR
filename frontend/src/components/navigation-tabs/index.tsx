import { Paper, Button } from '@mui/material';
import { useNavigate, useLocation } from 'react-router-dom';
import {
  Work as WorkIcon,
  People as PeopleIcon,
  EventNote as EventIcon,
} from '@mui/icons-material';

const NavigationButtons = () => {
  const navigate = useNavigate();
  const location = useLocation();

  const isActive = (path: string): boolean => {
    return location.pathname.startsWith(path);
  };

  const handleNavigation = (path: string) => {
    navigate(path);
  };

  const buttonStyle = {
    flex: 1,
    borderRadius: 2,
    py: 1.5,
    display: 'flex',
    flexDirection: 'column',
    gap: 0.5,
  };

  return (
    <Paper
      elevation={2}
      sx={{
        p: 2,
        mb: 3,
        display: 'flex',
        gap: 2,
        borderRadius: 3,
      }}
    >
      <Button
        variant={isActive('/vacancies') ? 'contained' : 'outlined'}
        onClick={() => handleNavigation('/vacancies')}
        sx={buttonStyle}
        startIcon={<WorkIcon />}
      >
        Вакансии
      </Button>

      <Button
        variant={isActive('/candidates') ? 'contained' : 'outlined'}
        onClick={() => handleNavigation('/candidates')}
        sx={buttonStyle}
        startIcon={<PeopleIcon />}
      >
        Кандидаты
      </Button>

      <Button
        variant={isActive('/interviews') ? 'contained' : 'outlined'}
        onClick={() => handleNavigation('/interviews')}
        sx={buttonStyle}
        startIcon={<EventIcon />}
      >
        Интервью
      </Button>
    </Paper>
  );
};

export default NavigationButtons;
