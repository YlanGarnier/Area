import { StyleSheet } from "react-native";

const styles = StyleSheet.create({
    centeredView: {
        flex: 1,
        justifyContent: "center",
        alignItems: "center",
        marginTop: 22,
    },
    modalView: {
        width: '80%',
        margin: 20,
        backgroundColor: "white",
        borderRadius: 20,
        padding: 35,
        alignItems: "center",
        shadowColor: "#000",
        shadowOffset: {
            width: 0,
            height: 2,
        },
        shadowOpacity: 0.25,
        shadowRadius: 4,
        elevation: 5,
    },
    modalOption: {
        padding: 10,
        borderBottomWidth: 1,
        borderColor: '#ddd',
    },
    modalText: {
        marginBottom: 15,
        textAlign: "center",
        color: '#725A8A',
    },
});

export default styles;
