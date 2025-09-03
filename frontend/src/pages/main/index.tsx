import { Container, Typography, Paper, Box } from '@mui/material';
import NavigationButtons from '../../components/navigation-tabs';

const MainPage = () => (
  <Container maxWidth='lg'>
    <Paper elevation={3} sx={{ p: 4, mt: 2, mb: 3 }}>
      <Box textAlign='center'>
        <Typography variant='h1' gutterBottom>
          AI-HR
        </Typography>
        <Typography variant='body1' gutterBottom>
          AI HR to automate the initial stages of the recruitment process
        </Typography>
      </Box>
    </Paper>
    <NavigationButtons />
  </Container>
);

export default MainPage;
