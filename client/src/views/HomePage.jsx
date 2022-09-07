import AddForm from "./AddForm/AddForm";
import ContactList from "./ContactList/ContactList";
import "./home-page.scss";

const HomePage = () => {
  return (
    <div className="home-page-container">
      <AddForm />
      <ContactList />
    </div>
  );
};

export default HomePage;
