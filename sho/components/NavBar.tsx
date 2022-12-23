import * as React from 'react';
import AppBar from '@mui/material/AppBar';
import Box from '@mui/material/Box';
import Toolbar from '@mui/material/Toolbar';
import Typography from '@mui/material/Typography';
import Button from '@mui/material/Button';

const pages = ['Cocktails', 'Beer', 'Wine', 'Coffee']
export default function NavBar() {
  return (
    <Box sx={{ flexGrow: 1}}>
      <AppBar position="static" sx={{height:'140px', background: 'white',}}>
        <Toolbar sx={{height:'100%', justifyContent:'space-between'}}>
          <Typography variant="h4" component="div" sx={{ flexGrow: 0, color: '#C51D1D', fontFamily: 'Jacques Francois', fontSize: '64px' }}>
            Sho!
          </Typography>

          <Box sx={{ flexGrow: 1, display: 'flex', mx:3}}>
            {pages.map((page) => (
              <Button
                key={page}
                sx={{ my: 2, paddingBottom: 0, marginBottom:0, color: 'black', fontFamily: 'Jacques Francois', display: 'block' }}
              >
                {page}
              </Button>
            ))}
          </Box>

        </Toolbar>
      </AppBar>
    </Box>
  );
}