import * as React from "react";

function MenuIcon(props: React.SVGProps<SVGSVGElement>) {
    return (
        <svg
            width={80}
            height={60}
            fill="none"
            xmlns="http://www.w3.org/2000/svg"
            {...props}
        >
            <path
                d="M0 0h72a8 8 0 018 8v44a8 8 0 01-8 8H0V0z"
                fill="#fff"
                fillOpacity={0.05}
            />
            <path
                d="M20 20h40M20 30h40M20 40h40"
                stroke="#F6F6F8"
                strokeWidth={3}
                strokeLinecap="round"
            />
        </svg>
    );
}

export default MenuIcon;