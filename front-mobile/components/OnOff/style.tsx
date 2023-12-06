import { StyleSheet } from "react-native";

const styles = StyleSheet.create({
    buttonContainer: {
        flex: 1,
        justifyContent: 'center',
        alignItems: 'center',
    },
    button: {
        width: 120,
        height: 40,
        borderRadius: 20,
        justifyContent: 'center',
        alignItems: 'center',
        padding: 3,
        borderColor: '#ccc',
        borderWidth: 2,
    },
    on: {
        backgroundColor: '#725A8A',
    },
    off: {
        backgroundColor: 'rgba(207, 147, 217, 1)',
    },
    slider: {
        width: 32,
        height: 32,
        borderRadius: 18,
        backgroundColor: 'white',
        position: 'absolute',
        top: 2,
    },
    sliderOn: {
        backgroundColor: 'rgba(143, 84, 149, 1)',
    },
    sliderOff: {
        backgroundColor: 'rgba(125, 114, 192, 1)',
    },
    label: {
        position: 'absolute',
        fontSize: 14,
        color: 'white',
        fontWeight: 'bold',
        marginTop: 3,
    },
});

export default styles;


