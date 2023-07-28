import React from 'react';

import styles from './button.module.css'
import { BUTTON_TYPES } from '@/common/data/button';

// TODO: Define a props
function Button(props: any) {
    const { type, btnText } = props;

    const getClassName = (type: string) => {
        switch (type) {
            case BUTTON_TYPES.PRIMARY:
                return `${styles.button} ${styles.primaryBtn}`;
                break;

            case BUTTON_TYPES.SECONDARY:
                return `${styles.button} ${styles.secondaryBtn}`;
                break;

            case BUTTON_TYPES.TERTIARY:
                return `${styles.button} ${styles.tertiaryBtn}`;
                break;

        }
    }
    return <div className={`${getClassName(type)}`}>{btnText}</div>;
}

export default Button;