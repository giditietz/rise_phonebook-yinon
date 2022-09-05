import { useState } from "react";
import TextBox from "../../components/Textbox";
import Title from "../../components/Title";

const AddForm = () => {
  const [firstName, setFirstName] = useState("");
  return (
    <div>
      <Title title={"First Name"} />
      <TextBox value={firstName} onChange={setFirstName} />
    </div>
  );
};

export default AddForm;
