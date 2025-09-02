import { Container, Typography, Paper, Box } from '@mui/material';

const MainPage = () => (
  <Container maxWidth='lg'>
    <Paper elevation={3} sx={{ p: 4, mt: 2 }}>
      <Box textAlign='center'>
        <Typography variant='h1' gutterBottom>
          AI-HR
        </Typography>
        <Typography variant='body1' gutterBottom>
          AI HR to automate the initial stages of the recruitment process
        </Typography>
      </Box>
    </Paper>
  </Container>
);

export default MainPage;
