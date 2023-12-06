import * as React from "react"
import { Svg, Path } from 'react-native-svg';
const LogoSlack = (props) => (
    <Svg
        xmlns="http://www.w3.org/2000/svg"
        width={100}
        height={100}
        viewBox="0 0 48 48"
        {...props}
    >
        <Path
            fill="#f5bc00"
            d="M25 29c0 2.2 1.8 4 4 4h10c2.2 0 4-1.8 4-4s-1.8-4-4-4H29c-2.2 0-4 1.8-4 4zM25 35v4c0 2.2 1.8 4 4 4s4-1.8 4-4-1.8-4-4-4h-4z"
        />
        <Path
            fill="#f55376"
            d="M19 25c-2.2 0-4 1.8-4 4v10c0 2.2 1.8 4 4 4s4-1.8 4-4V29c0-2.2-1.8-4-4-4zM13 25H9c-2.2 0-4 1.8-4 4s1.8 4 4 4 4-1.8 4-4v-4z"
        />
        <Path
            fill="#00b3d7"
            d="M23 19c0-2.2-1.8-4-4-4H9c-2.2 0-4 1.8-4 4s1.8 4 4 4h10c2.2 0 4-1.8 4-4zM23 13V9c0-2.2-1.8-4-4-4s-4 1.8-4 4 1.8 4 4 4h4z"
        />
        <Path
            fill="#00b569"
            d="M29 23c2.2 0 4-1.8 4-4V9c0-2.2-1.8-4-4-4s-4 1.8-4 4v10c0 2.2 1.8 4 4 4zM35 23h4c2.2 0 4-1.8 4-4s-1.8-4-4-4-4 1.8-4 4v4z"
        />
    </Svg>
)
export default LogoSlack
