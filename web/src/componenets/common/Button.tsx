import * as React from 'react';

import Button, { ButtonProps } from '@mui/material/Button';
import { styled } from '@mui/material/styles';

const PrimaryButton = styled(Button)(({ theme }) =>({ 
  background: '#161B33',
  color: '#F1DAC4',
  width: '20vw'
 })); 
  
const SecondaryButton = styled(Button)({  
  background: '#F1DAC4',
  color: '#161B33',
  width: '20vw',
});

export default function StyledButton(props: any) {
  if (props.variant == 'primary') {
    return (
      <PrimaryButton className={props.className}>{props.children}</PrimaryButton>
    )
  } 
  return (
    <SecondaryButton className={props.className}>{props.children}</SecondaryButton>
  );
}