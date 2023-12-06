import * as React from "react"
import { Svg, Path } from 'react-native-svg';
const LogoDiscord = (props) => (
    <Svg
        xmlns="http://www.w3.org/2000/svg"
        width={100}
        height={100}
        viewBox="0 0 48 48"
        {...props}
    >
        <Path fill="#7880e7" d="M11 24 25 2l14 22-14 8-14-8z" />
        <Path fill="#5c64c7" d="m25 2 14 22-14 8V2z" />
        <Path fill="#7880e7" d="m11 27 14 8 14-8-14 19-14-19z" />
        <Path fill="#5c64c7" d="m25 35 14-8-14 19V35zM11 24l14-6 14 6-14 8-14-8z" />
        <Path fill="#2a3192" d="m25 18 14 6-14 8V18z" />
    </Svg>
)
export default LogoDiscord
