import Upload from "../Upload";
import Item from "./Item";
import Image from "next/image";
import Link from "next/link"

const SiteNav = () => {
    return (
        <nav className="absolute bottom-0 p-2 w-full bg-gray-100 grid grid-cols-4 justify-center text-black border-t border-gray-300 md:static md:h-screen md:row-start-1 md:row-span-full md:col-start-1 md:col-end-4 md:flex md:flex-col md:justify-start md:space-y-2 lg:col-end-3">
            <Link href="/" className="hidden md:flex md:place-items-center md:p-3">
                <Image
                    className="md:col-span-2 md:rounded-full md:my-2"
                    src="/logo.png"
                    alt="Vaulter Logo"
                    width={40}
                    height={40}
                ></Image>
                <span className="md:ml-2 md:text-2xl">Vaulter</span>
            </Link>
            <Upload />
            <Item src="/icons/files.svg" alt="Files Icon" name="Files" />
            <Item src="/icons/starred.svg" alt="Starred Icon" name="Starred" />
            <Item src="/icons/recent.svg" alt="Recent Icon" name="Recent" />
            <Item src="/icons/profile.svg" alt="Profile Icon" name="Profile" />
        </nav>
    );
};

export default SiteNav;
