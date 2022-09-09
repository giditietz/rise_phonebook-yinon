import { useState } from "react";
import SaveIcon from "@mui/icons-material/Save";
import Button from "@mui/material/Button";
import FormField from "../../components/FormField";

import text from "../../utils/language/text.json";

import "./phone-row.scss";

const PhoneRow = ({
  descriptionValue,
  phoneNumberValue,
  onPhoneSave,
  index,
}) => {
  const [description, setDescription] = useState(descriptionValue);
  const [phoneNumber, setPhoneNumber] = useState(phoneNumberValue);

  const onSave = () => {
    onPhoneSave(description, phoneNumber, index);
  };

  return (
    <div className="phone-container">
      <FormField
        title={text.description}
        value={description}
        onChange={setDescription}
      />
      <FormField
        title={text.phoneNumber}
        value={phoneNumber}
        onChange={setPhoneNumber}
      />
      <Button
        onClick={onSave}
        style={{ textTransform: "none" }}
        variant="contained"
        startIcon={<SaveIcon />}
      >
        {text.save}
      </Button>
    </div>
  );
};

export default PhoneRow;
