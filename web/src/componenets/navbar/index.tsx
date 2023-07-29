import React from 'react';
import style from './navbar.module.css';

// TODO: Define a props
function NavBar(props: any) {
    const { name } = props;

    return (
        <div className={style.navBar}>
            <a href='/'>
                <div className={style.siteName}>Jeffy's Bar</div>
            </a>
            
            <ul>
                <li className={style.loginBtn}>
                    <a href="/login">Login</a>
                </li>
            </ul>
            
        </div>
    );


}

export default NavBar;