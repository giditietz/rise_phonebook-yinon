import { useState } from "react";

import FormField from "../../components/FormField";
import text from "../../utils/language/text.json";

import "./add-form.scss";
import AddressRow from "./AddressRow";
import PhoneRow from "./PhoneRow";

const AddForm = () => {
  const [firstName, setFirstName] = useState("");
  const [lastName, setLastName] = useState("");
  const [phones, setPhones] = useState([{ description: "", phoneNumber: "" }]);
  const [addresses, setAddresses] = useState([
    { description: "", city: "", street: "", homeNumber: "", apartment: "" },
  ]);

  const onPhoneSave = (description, phoneNumber, index) => {
    const newPhone = { ...phones[index] };
    newPhone.description = description;
    newPhone.phoneNumber = phoneNumber;
    const newPhoneArray = [
      ...phones.slice(0, index),
      newPhone,
      ...phones.slice(index + 1),
    ];
    if (newPhoneArray[newPhoneArray.length - 1].phoneNumber !== "") {
      newPhoneArray.push({ description: "", phoneNumber: "" });
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
    newAddress.homeNumber = homeNumber;
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
        homeNumber: "",
        apartment: "",
      });
    }
    setAddresses(newAddressArray);
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
      <h1 className="add-content">{text.addPhones}</h1>
      {phones.map((phone, index) => (
        <PhoneRow
          key={`${phone.description}${index}`}
          descriptionValue={phone.description}
          phoneNumberValue={phone.phoneNumber}
          onPhoneSave={onPhoneSave}
          index={index}
        />
      ))}
      <h1 className="add-content">{text.addAddress}</h1>
      {addresses.map((address, index) => (
        <AddressRow
          key={`${address.description}${index}`}
          descriptionValue={address.description}
          cityValue={address.city}
          streetValue={address.street}
          homeNumberValue={address.homeNumber}
          apartmentValue={address.apartment}
          index={index}
          onAddressSave={onAddressSave}
        />
      ))}
    </div>
  );
};

export default AddForm;
