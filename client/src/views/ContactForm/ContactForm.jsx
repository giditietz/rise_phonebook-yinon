import { useEffect, useState } from "react";

import SaveIcon from "@mui/icons-material/Save";
import Button from "@mui/material/Button";

import FormField from "../../components/FormField";
import AddressRow from "./AddressRow/AddressRow";
import PhoneRow from "./PhoneRow/PhoneRow";

import httpRequest from "../../utils/httpRequest/httpRequest";

import text from "../../utils/language/text.json";
import "./contact-form.scss";

const ContactForm = ({ contact, isShow, isNew, isEdit, handleSubmit }) => {
  const [firstName, setFirstName] = useState(isNew ? "" : contact?.firstName);
  const [lastName, setLastName] = useState(isNew ? "" : contact?.lastName);
  const [phones, setPhones] = useState(
    isNew ? [{ description: "", phone_number: "" }] : contact?.phone
  );
  const [addresses, setAddresses] = useState(
    isNew
      ? [
          {
            description: "",
            city: "",
            street: "",
            home_number: "",
            apartment: "",
          },
        ]
      : contact?.address
  );

  const setEditArray = () => {
    if (isEdit) {
      let editAddresses = [];
      if (addresses) {
        editAddresses = [...addresses];
      }
      editAddresses.push({
        description: "",
        city: "",
        street: "",
        home_number: "",
        apartment: "",
      });
      setAddresses(editAddresses);
      let editPhones = [];
      if (phones) {
        editPhones = [...phones];
      }
      editPhones.push({ description: "", phone_number: "" });
      setPhones(editPhones);
    }
  };

  useEffect(() => {
    setEditArray();
  }, [contact]);

  const onPhoneSave = (description, phoneNumber, index) => {
    const newPhone = { ...phones[index] };
    newPhone.description = description;
    newPhone.phone_number = phoneNumber;
    const newPhoneArray = [
      ...phones.slice(0, index),
      newPhone,
      ...phones.slice(index + 1),
    ];
    if (newPhoneArray[newPhoneArray.length - 1].phone_number !== "") {
      newPhoneArray.push({ description: "", phone_number: "" });
    }
    setPhones(newPhoneArray);
  };

  const onAddressSave = (
    description,
    city,
    street,
    homeNumber,
    apartment,
    index
  ) => {
    const newAddress = { ...addresses[index] };
    newAddress.description = description;
    newAddress.city = city;
    newAddress.street = street;
    newAddress.home_number = homeNumber;
    newAddress.apartment = apartment;
    const newAddressArray = [
      ...addresses.slice(0, index),
      newAddress,
      ...addresses.slice(index + 1),
    ];
    if (newAddressArray[newAddressArray.length - 1].city !== "") {
      newAddressArray.push({
        description: "",
        city: "",
        street: "",
        home_number: "",
        apartment: "",
      });
    }
    setAddresses(newAddressArray);
  };

  const onSubmit = () => {
    if (phones[phones.length - 1].phone_number === "") {
      phones.pop();
    }
    if (addresses[addresses.length - 1].city === "") {
      addresses.pop();
    }

    const newContact = {
      first_name: firstName,
      last_name: lastName,
      address: addresses,
      phone: phones,
    };
    if (isEdit) {
      httpRequest.put(`/contacts/${contact.contactID}`, newContact);
    } else {
      httpRequest.post("/contacts", newContact);
    }
    handleSubmit();
  };

  return (
    <div className="form-container">
      <div className="form-field">
        <FormField
          title={text.firstName}
          value={firstName}
          onChange={setFirstName}
        />
      </div>
      <div className="form-field">
        <FormField
          title={text.lastName}
          value={lastName}
          onChange={setLastName}
        />
      </div>
      {phones?.length ? (
        isShow ? (
          <h1 className="add-content">{text.phones}</h1>
        ) : (
          <h1 className="add-content">{text.addPhones}</h1>
        )
      ) : null}
      {phones?.length
        ? phones?.map((phone, index) => (
            <PhoneRow
              key={`${phone.description}${index}`}
              descriptionValue={phone.description}
              phoneNumberValue={phone.phone_number}
              onPhoneSave={onPhoneSave}
              index={index}
              isSave={!isShow}
            />
          ))
        : null}
      {addresses?.length ? (
        isShow ? (
          <h1 className="add-content">{text.addresses}</h1>
        ) : (
          <h1 className="add-content">{text.addAddress}</h1>
        )
      ) : null}
      {addresses?.length
        ? addresses.map((address, index) => (
            <AddressRow
              isSave={!isShow}
              key={`${address.description}${index}`}
              descriptionValue={address.description}
              cityValue={address.city}
              streetValue={address.street}
              homeNumberValue={address.home_number}
              apartmentValue={address.apartment}
              index={index}
              onAddressSave={onAddressSave}
            />
          ))
        : null}
      {!isShow ? (
        <Button
          onClick={onSubmit}
          style={{ textTransform: "none", width: "60%", left: "20%" }}
          variant="contained"
          startIcon={<SaveIcon />}
        >
          {text.submit}
        </Button>
      ) : null}
    </div>
  );
};

export default ContactForm;
