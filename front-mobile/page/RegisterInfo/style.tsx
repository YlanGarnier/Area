import { StyleSheet } from 'react-native';

const styles = StyleSheet.create({
    container: {
        flex: 1,
        backgroundColor: '#D0CBED',
        alignItems: 'center',
        justifyContent: 'center',
    },
    background: {
        width: '100%',
        height: '30%',
        position: 'absolute',
        top: 0,
    },
    catchphrase: {
        marginVertical: 30,
        fontSize: 20,
        fontWeight: 'bold',
        color: '#f94370',
        marginTop: '50%',
    },
    middle: {
        width: '85%',
        marginTop: '10%',
        textAlign: 'center',
        justifyContent: 'center',
        alignItems: 'center',
    },
    label: {
        marginBottom: 5,
        paddingLeft: 10,
        fontSize: 17,
        color: '#f94370',
        fontWeight: 'bold',
    },
    inputContainer: {
        textAlign: 'center',
        alignItems: 'center',
        flexDirection: 'column',
        width: '100%',
        height: 60,
        marginBottom: 20,
        justifyContent: 'center',
    },
    input: {
        width: 300,
        height: 40,
        borderWidth: 1,
        borderRadius: 15,
        paddingLeft: 5,
        marginBottom: 10,
        backgroundColor: 'rgba(255, 255, 255, 0.27)',
    },
    signButton: {
        alignItems: 'center',
        justifyContent: 'center',
        paddingVertical: 12,
        paddingHorizontal: 50,
        borderRadius: 25,
        marginTop: 10,
        backgroundColor: '#f94370',
        elevation: 5,
        shadowColor: '#000',
        shadowOffset: { width: 0, height: 2 },
        shadowOpacity: 0.3,
        shadowRadius: 3,
    },
    signinText: {
        fontSize: 20,
        fontWeight: 'bold',
        textAlign: 'center',
        color: '#FFFFFF',
    },
    messageContainer: {
        width: '100%',
        alignItems: 'center',
        justifyContent: 'center',
        marginVertical: 10,
    },
    errmessage: {
        color: "red",
        textAlign: 'center',
        width: '100%',
    },
});

export default styles;