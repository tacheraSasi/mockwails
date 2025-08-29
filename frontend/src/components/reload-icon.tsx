import { LoaderCircleIcon, RefreshCcw } from "lucide-react";
import { useState } from "react";

const ReloadIcon = () => {
  const [spinning, setSpinning] = useState(false);

  const handleClick = () => {
    if (spinning) return;
    setSpinning(true);
    setTimeout(() => setSpinning(false), 1000); // 1s spin
  };

  return (
    <button
      type="button"
      aria-label="Reload"
      onClick={handleClick}
      className="focus:outline-none"
      style={{
        display: "flex",
        alignItems: "center",
        justifyContent: "center",
        background: "none",
        border: "none",
        padding: 0,
        cursor: "pointer",
      }}
    >
      <RefreshCcw 
        className={
          spinning ? "animate-spin-smooth text-primary" : "text-primary"
        }
        size={24}
        style={{ transition: "color 0.2s" }}
      />
    </button>
  );
};

export default ReloadIcon;
