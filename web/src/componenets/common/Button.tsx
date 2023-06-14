import * as React from 'react';

import Button, { ButtonProps } from '@mui/material/Button';
import { styled } from '@mui/material/styles';

const PrimaryButton = styled(Button)({ 
  background: '#161B33',
  color: '#F1DAC4',
 }); 
  
const SecondaryButton = styled(Button)({  
  background: '#F1DAC4',
  color: '#161B33',
});

export default function StyledButton(props: any) {
  if (props.variant == 'primary') {
    return (
      <PrimaryButton>{props.children}</PrimaryButton>
    )
  } 
  return (
    <SecondaryButton>{props.children}</SecondaryButton>
  );
}