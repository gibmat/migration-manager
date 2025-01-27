import { FC, useState } from 'react';
import ReactDOM from 'react-dom';
import { MdOutlineComment } from 'react-icons/md';
import { Instance } from 'types/instance';

interface Props {
  instance: Instance;
}

interface MousePosition {
  top: number;
  left: number;
}

const InstanceActions: FC<Props> = ({instance}) => {
  const [tooltipPosition, setTooltipPosition] = useState<MousePosition>({top: 0, left: 0});

  const handleMouseEnter = (e: any) => {
    const rect = e.target.getBoundingClientRect();
    setTooltipPosition({
      top: rect.top,
      left: rect.left - 100,
    });
  };

  const handleMouseLeave = () => {
    setTooltipPosition({top: 0, left: 0}); // Hide the tooltip
  };

  return (
    <div className="relative inline-block">
      <div>
        { instance.overrides && instance.overrides.comment && (
          <MdOutlineComment
            size={15}
            onMouseEnter={handleMouseEnter}
            onMouseLeave={handleMouseLeave}
            />
        )}
      </div>
      {tooltipPosition.top != 0 &&
        ReactDOM.createPortal(
          <div
            style={{
              position: "absolute",
              top: tooltipPosition.top,
              left: tooltipPosition.left,
              transform: "translate(-50%, 0%)",
              padding: "8px",
              background: "rgba(0, 0, 0, 0.8)",
              color: "white",
              borderRadius: "4px",
              pointerEvents: "none",
              zIndex: 1000,
            }}
          >
            {instance.overrides.comment}
          </div>,
          document.body
        )}
    </div>
  );
};

export default InstanceActions;
