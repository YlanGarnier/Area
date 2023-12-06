import * as React from "react"
import { Svg, Path, LinearGradient, Stop } from 'react-native-svg';
const LogoMiro = (props) => (
    <Svg
        xmlns="http://www.w3.org/2000/svg"
        width={props.width || 50}
        height={props.height || 50}
        viewBox="0 0 64 64"
        {...props}
    >
        <Path
            fill="none"
            stroke="url(#1dP_ejbheoOkMyPxlPY0Wa)"
            strokeMiterlimit={10}
            strokeWidth={2}
            d="M18.4 8h27.2C51.344 8 56 12.656 56 18.4v27.2C56 51.344 51.344 56 45.6 56H18.4C12.656 56 8 51.344 8 45.6V18.4C8 12.656 12.656 8 18.4 8z"
        />
        <LinearGradient
            id="a"
            x1={33}
            x2={33}
            y1={50}
            y2={14}
            gradientUnits="userSpaceOnUse"
        >
            <Stop offset={0} stopColor="#e6abff" />
            <Stop offset={1} stopColor="#6dc7ff" />
        </LinearGradient>
        <Path
            fill="url(#a)"
            fillRule="evenodd"
            d="M41.341 14h-5.268l4.39 7.714L30.805 14h-5.268l4.829 9.429L20.268 14H15l5.268 12L15 50h5.268l10.098-25.714L25.537 50h5.268l9.659-27.429L36.073 50h5.268L51 20l-9.659-6z"
            clipRule="evenodd"
        />
    </Svg>
)
export default LogoMiro
