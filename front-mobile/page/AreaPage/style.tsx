import { StyleSheet } from "react-native";

const styles = StyleSheet.create({
  container: {
    flex: 1,
    backgroundColor: "#D0CBED",
    alignItems: "center",
    paddingHorizontal: 20,
  },
  contentContainer: {
    flex: 1,
    flexDirection: "column",
    justifyContent: "space-evenly",
    alignItems: "center",
    width: "100%",
  },
  appZone: {
    backgroundColor: "white",
    alignItems: "center",
    justifyContent: "center",
    width: "90%",
    height: "30%",
    borderRadius: 10,
  },
  appAction: {
    fontSize: 28,
    fontWeight: "bold",
    marginTop: 10,
  },
  verticalLine: {
    width: 5,
    backgroundColor: "#7D72C0",
    fontWeight: "bold",
    height: "10%",
  },
  text: {
    fontSize: 22,
    marginBottom: 20,
    color: "#7D72C0",
  },
  deleteButton: {
    marginBottom: 10,
  },
  confirmationModal: {
    backgroundColor: "white",
    padding: 20,
    borderRadius: 10,
  },
  confirmationText: {
    fontSize: 18,
    marginBottom: 20,
  },
  confirmationButtons: {
    flexDirection: "row",
    justifyContent: "space-around",
  },
  confirmationButton: {
    backgroundColor: "#8F5495",
    padding: 10,
    borderRadius: 5,
  },
  confirmationButtonText: {
    color: "white",
    fontSize: 16,
  },
});

export default styles;
