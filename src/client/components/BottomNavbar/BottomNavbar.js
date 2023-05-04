import Item from "./Item";

const BottomNavbar = () => {
    return (
        <nav className="absolute bottom-0 p-1 bg-gray-100 grid grid-cols-3 w-full justify-center text-black border-t border-gray-300">
            <Item src="/icons/files.svg" alt="Files Icon" name="Files" />
            <Item src="/icons/starred.svg" alt="Starred Icon" name="Starred" />
            <Item src="/icons/profile.svg" alt="Profile Icon" name="Profile" />
        </nav>
    );
};

export default BottomNavbar;
