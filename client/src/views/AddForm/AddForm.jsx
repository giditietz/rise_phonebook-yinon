import { useState } from "react";
import httpRequest from "../../utils/httpRequest/httpRequest";
import ContactFields from "./ContactField";

const AddForm = () => {
  const [firstName, setFirstName] = useState("");
  const [lastName, setLastName] = useState("");

  return (
    <ContactFields
      firstName={firstName}
      setFirstName={setFirstName}
      lastName={lastName}
      setLastName={setLastName}
    />
  );
};

export default AddForm;
