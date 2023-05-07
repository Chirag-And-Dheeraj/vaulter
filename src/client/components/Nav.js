import Image from "next/image";

const Nav = () => {
    return (
            <nav className="grid grid-cols-12 p-3">
                <Image
                    className="col-span-2 sm:col-span-1 rounded-full md:hidden"
                    src="/logo.png"
                    alt="Vaulter Logo"
                    width={40}
                    height={40}
                ></Image>
                <div className="col-start-3 col-end-13 rounded-md border-2 border-gray-200 sm:col-start-2 md:col-span-full md:place-self-start md:justify-self-stretch">
                    <label htmlFor="search" className="sr-only">Search by folder or file name</label>
                    <input
                        id="search"
                        className="w-full p-2"
                        type="text"
                        placeholder="Search by folder or file name"
                    />
                </div>
            </nav>
    );
};

export default Nav;
