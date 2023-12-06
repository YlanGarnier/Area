import * as React from "react"
import { Svg, Path } from 'react-native-svg';
const LogoTwitch = (props) => (
    <Svg
        xmlns="http://www.w3.org/2000/svg"
        width={props.width || 50}
        height={props.height || 50}
        viewBox="0 0 48 48"
        {...props}
    >
        <Path
            fill="#7e57c2"
            d="M42 27.676 33 38h-7l-7 6h-3v-6H6V15.706L10.364 7H42v20.676z"
        />
        <Path
            fill="#fafafa"
            d="M39 26.369 34 32h-7l-7 6v-6l-8-.024V10h27v16.369z"
        />
        <Path fill="#7e57c2" d="M21 16h3v10h-3zM30 16h3v10h-3z" />
    </Svg>
)
export default LogoTwitch
