import Toast from 'react-bootstrap/Toast';
import ToastContainer from 'react-bootstrap/ToastContainer';
import { MdOutlineClose } from "react-icons/md";
import { useNotification } from "context/notification";

const Notification = () => {
  const { notify, notification } = useNotification();

  const closeToast = () => {
    notify.success("");
  };

  return (
    <>
    {notification.message != "" && (
    <ToastContainer
      className="p-3"
      position="bottom-end"
      style={{ zIndex: 1 }}
    >
      <Toast bg={notification.type} >
        <Toast.Body className="text-white">
          <div className="container">
            <div><p className="float-end"><MdOutlineClose style={{cursor: 'pointer'}} onClick={() => closeToast()} /></p></div>
            <div>{ notification.message }</div>
          </div>
        </Toast.Body>
      </Toast>
    </ToastContainer>
    )}
    </>
  );
};

export default Notification;

