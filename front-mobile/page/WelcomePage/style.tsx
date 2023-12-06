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
        height: '40%',
        position: 'absolute',
        top: 0,
    },
    middle: {
        marginTop: '75%',
        textAlign: 'center',
        justifyContent: 'center',
    },
    catchphrase: {
        margin: 20,
        fontSize: 20,
        fontWeight: 'bold',
        color: '#7D72C0',
    },
    inputContainer: {
        flexDirection: 'column',
        justifyContent: 'flex-start',
        alignItems: 'center',
    },
    label: {
        marginBottom: 5,
        paddingLeft: 10,
        fontSize: 17,
        color: '#C3A6E5',
        fontWeight: 'bold',
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
        backgroundColor: '#8F5495',
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
    registerContainer: {
        flexDirection: 'row',
        alignItems: 'center',
        marginTop: 50,
    },
    registerText: {
        fontSize: 16,
        color: '#7D72C0',
    },
    signupText: {
        fontSize: 17,
        fontWeight: 'bold',
        color: '#725A8A',
    },

    socialContainer: {
        flexDirection: 'row',
    },
    social: {
        marginTop: 10,
        marginHorizontal: 30,
    },
    icon: {
        width: 40,
        height: 40,
    },
    dividerContainer: {
        flexDirection: 'row',
        alignItems: 'center',
        marginVertical: 20,
        width: '80%',
    },

    dividerLine: {
        flex: 1,
        height: 1,
        backgroundColor: '#C3A6E5',
    },

    dividerText: {
        marginHorizontal: 10,
        color: '#8F5495',
    },
    messageContainer: {
        display: "flex",
        alignItems: "center",
        justifyContent: "center"
    },
    errmessage: {
        color: "red"
    },
});

export default styles;
