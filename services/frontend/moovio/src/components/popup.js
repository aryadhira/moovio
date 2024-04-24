import React from 'react';

const PopUp = ({ isOpen, onClose, children }) => {
    if (!isOpen) return null;

    return (
        <div className="fixed inset-0 flex items-center justify-center">
            <div className="bg-gray-800 bg-opacity-50 absolute inset-0"></div>
            <div className="bg-white p-8 rounded-lg shadow-lg z-10">
                <button className="absolute top-2 right-2 text-gray-500 hover:text-gray-800" onClick={onClose}>
                    <svg xmlns="http://www.w3.org/2000/svg" className="h-6 w-6" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                        <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M6 18L18 6M6 6l12 12" />
                    </svg>
                </button>
                {children}
            </div>
        </div>
    );
};

export default PopUp;
