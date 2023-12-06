import { StyleSheet } from "react-native";

const styles = StyleSheet.create({
    container: {
        flex: 1,
        alignItems: 'center',
        padding: 20,
    },
    detailsCard: {
        backgroundColor: 'rgba(255, 255, 255, 1)',
        padding: 20,
        borderRadius: 20,
        width: '95%',
        marginVertical: 30,
        elevation: 5,
        shadowColor: "#000",
        shadowOffset: {
            width: 0,
            height: 4,
        },
        shadowOpacity: 0.32,
        shadowRadius: 5.46,
    },
    nameText: {
        fontSize: 24,
        fontWeight: 'bold',
        marginBottom: 15,
        color: '#4B3F72',
    },
    detailsText: {
        flexDirection: 'row',
        alignItems: 'center',
        fontSize: 18,
        marginBottom: 15,
        color: '#4B3F72',
    },
    SignOutText: {
        color: '#137c8b',
        fontSize: 16,
        textAlign: 'center',
        margin: 10
    },
    DeleteAccountText: {
        color: '#D32F2F',
        fontSize: 16,
        textAlign: 'center',
        margin: 10
    },
    bottomModalContainer: {
        flex: 1,
        justifyContent: 'flex-end',
        alignItems: 'center',
    },
    modalView: {
        width: '100%',
        backgroundColor: 'white',
        borderTopLeftRadius: 20,
        borderTopRightRadius: 20,
        padding: 15,
        alignItems: 'center',
    },
    modalText: {
        marginBottom: 15,
        textAlign: 'center',
        fontSize: 18,
        fontWeight: 'bold'
    },
    modalButton: {
        backgroundColor: 'transparent',
        padding: 10,
        margin: 5,
        borderRadius: 0,
    },
    modalCancelButton: {
        backgroundColor: 'transparent',
        padding: 10,
        margin: 5,
        borderRadius: 0,
    },
    linkText: {
        color: '#137c8b',
        fontSize: 16,
        textAlign: 'center',
        margin: 10
    },
    cancelText: {
        color: '#999',
        fontSize: 16,
        textAlign: 'center',
        margin: 10
    },
    editButton: {
        backgroundColor: '#4B3F72',
        alignItems: 'center',
        justifyContent: 'center',
        borderRadius: 15,
        height: 40,
        width: '30%',
    },
    editText: {
        color: 'white',
        fontSize: 14,
        textAlign: 'center',
        margin: 10,
        fontWeight: 'bold'
    },
    modalContainer: {
        flex: 1,
        padding: 20,
        marginTop: "10%",
        backgroundColor: 'white'
    },
    editProfileHeader: {
        textAlign: 'center',
        fontSize: 30,
        fontWeight: 'bold',
        marginBottom: 50,
        color: '#725A8A',
    },
    inputContainer: {
        marginBottom: 15
    },
    inputLabel: {
        color: '#8F5495',
        fontSize: 16,
        fontWeight: 'bold',
        marginBottom: 5
    },
    input: {
        borderWidth: 2,
        borderColor: '#eee',
        padding: 10,
        borderRadius: 5,
        color: '#4B3F72',
    },
    editButtonsContainer: {
        flexDirection: 'row',
        justifyContent: 'space-between',
        marginBottom: 20,
    },
    saveText: {
        color: '#4B3F72',
        fontSize: 16,
        textAlign: 'center',
        margin: 10,
        fontWeight: 'bold'
    },
    closeText: {
        color: '#999',
        fontSize: 16,
        textAlign: 'center',
        margin: 10,
        fontWeight: 'bold'
    }
});

export default styles;
