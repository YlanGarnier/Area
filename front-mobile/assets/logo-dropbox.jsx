import * as React from "react"
import { Svg, Path } from 'react-native-svg';
const LogoDropbox = (props) => (
    <Svg
        xmlns="http://www.w3.org/2000/svg"
        width={100}
        height={100}
        viewBox="0 0 48 48"
        {...props}
    >
        <Path
            fill="#1E88E5"
            d="M42 13.976 31.377 7.255 24 13.314l11.026 6.418zM6 25.647l10.933 6.408L24 26.633l-10.472-6.664zM16.933 7.255 6 14.301l7.528 5.668L24 13.314zM24 26.633l7.209 5.422L42 25.647l-6.974-5.915z"
        />
        <Path
            fill="#1E88E5"
            d="m32.195 33.779-1.148.683-1.068-.804L24 29.162l-5.845 4.484-1.064.818-1.158-.679L13 32.066v2.672L23.988 42 35 34.794v-2.68z"
        />
    </Svg>
)
export default LogoDropbox
