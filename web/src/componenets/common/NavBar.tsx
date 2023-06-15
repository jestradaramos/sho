import * as React from 'react';
import AppBar from '@mui/material/AppBar';
import Box from '@mui/material/Box';
import Toolbar from '@mui/material/Toolbar';
import Typography from '@mui/material/Typography';
import Button from './Button';

export default function ButtonAppBar() {
  return (
    <Box sx={{ flexGrow: 1 }}>
      <AppBar style={{ background: 'transparent', boxShadow: 'none'}} position="static">
        <Toolbar>
          <Typography variant="h4" component="div" sx={{ flexGrow: 1 }} color='#F1DAC4'>
            Jeffy's Bar
          </Typography>
          <Button variant='primary' className='NavButton'>Login</Button>
        </Toolbar>
      </AppBar>
    </Box>
  );
}
