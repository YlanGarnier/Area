import * as React from "react"
import { Svg, Path } from 'react-native-svg';

const LogoHttp = (props) => (
    <Svg
        xmlns="http://www.w3.org/2000/svg"
        width={800}
        height={800}
        viewBox="0 0 32 32"
        {...props}
    >
        <Path d="M30 11h-5v10h2v-3h3a2.003 2.003 0 0 0 2-2v-3a2.002 2.002 0 0 0-2-2Zm-3 5v-3h3l.001 3ZM10 13h2v8h2v-8h2v-2h-6v2zM23 11h-6v2h2v8h2v-8h2v-2zM6 11v4H3v-4H1v10h2v-4h3v4h2V11H6z" />
        <Path
            d="M0 0h32v32H0z"
            data-name="&lt;Transparent Rectangle&gt;"
            style={{
                fill: "none",
            }}
        />
    </Svg>
)
export default LogoHttp
