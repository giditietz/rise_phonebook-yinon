import { useState } from "react";
import SaveIcon from "@mui/icons-material/Save";
import DeleteIcon from "@mui/icons-material/Delete";
import Button from "@mui/material/Button";
import FormField from "../../../components/FormField";

import httpRequest from "../../../utils/httpRequest/httpRequest";

import text from "../../../utils/language/text.json";

import "./phone-row.scss";

const PhoneRow = ({
  descriptionValue,
  phoneNumberValue,
  onPhoneSave,
  isSave,
  onPhoneDelete,
  phoneID,
  index,
}) => {
  const [description, setDescription] = useState(descriptionValue);
  const [phoneNumber, setPhoneNumber] = useState(phoneNumberValue);

  const onSave = () => {
    onPhoneSave(description, phoneNumber, index);
  };

  const onDelete = () => {
    if (phoneID) {
      httpRequest.del(`/contacts/phone/${phoneID}`);
    }
    onPhoneDelete(phoneID);
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
      {isSave ? (
        <Button
          onClick={onSave}
          style={{ textTransform: "none" }}
          variant="contained"
          startIcon={<SaveIcon />}
        >
          {text.save}
        </Button>
      ) : null}
      {isSave && phoneNumber !== "" ? (
        <Button
          onClick={onDelete}
          style={{ textTransform: "none" }}
          variant="contained"
          color="error"
          startIcon={<DeleteIcon />}
        >
          {text.delete}
        </Button>
      ) : null}
    </div>
  );
};

export default PhoneRow;
