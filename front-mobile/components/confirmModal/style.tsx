import { StyleSheet } from "react-native";

const styles = StyleSheet.create({
  container: {
    flex: 1,
    justifyContent: "center",
    alignItems: "center",
    backgroundColor: "rgba(0, 0, 0, 0.5)",
  },
  modal: {
    backgroundColor: "white",
    borderRadius: 10,
    padding: 20,
    width: 300,
  },
  text: {
    fontSize: 18,
    marginBottom: 20,
    textAlign: "center",
  },
  confirmButton: {
    fontSize: 18,
    color: "red",
    textAlign: "center",
    marginBottom: 10,
  },
});

export default styles;
