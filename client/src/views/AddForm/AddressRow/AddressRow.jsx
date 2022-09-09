import { useState } from "react";
import SaveIcon from "@mui/icons-material/Save";
import Button from "@mui/material/Button";
import FormField from "../../../components/FormField";

import text from "../../../utils/language/text.json";

import "./address-row.scss";

const AddressRow = ({
  descriptionValue,
  cityValue,
  streetValue,
  homeNumberValue,
  apartmentValue,
  index,
  onAddressSave,
}) => {
  const [description, setDescription] = useState(descriptionValue);
  const [city, setCity] = useState(cityValue);
  const [street, setStreet] = useState(streetValue);
  const [homeNumber, setHomeNumber] = useState(homeNumberValue);
  const [apartment, setApartment] = useState(apartmentValue);

  const onSave = () => {
    onAddressSave(description, city, street, homeNumber, apartment, index);
  };

  return (
    <>
      {" "}
      <div className="address-container">
        <FormField
          title={text.description}
          value={description}
          onChange={setDescription}
        />
        <FormField title={text.city} value={city} onChange={setCity} />
        <FormField title={text.street} value={street} onChange={setStreet} />
      </div>
      <div className="address-container">
        <FormField
          title={text.homeNumber}
          value={homeNumber}
          onChange={setHomeNumber}
        />
        <FormField
          title={text.apartment}
          value={apartment}
          onChange={setApartment}
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
    </>
  );
};

export default AddressRow;
