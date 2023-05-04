import Image from "next/image";

const Nav = () => {
    return (
        <>
            <nav className="grid grid-cols-12 items-center p-3">
                <Image
                    className="col-span-2 rounded-full"
                    src="/logo.png"
                    alt="Vaulter Logo"
                    width={40}
                    height={40}
                ></Image>
                <div className="col-span-10 grid grid-cols-1 items-centerborder-2 border-black">
                    <label htmlFor="search" className="sr-only"></label>
                    <input
                        id="search"
                        className="w-full p-2"
                        type="text"
                        placeholder="Search by folder or file name"
                    />
                </div>
            </nav>
            <hr />
        </>
    );
};

export default Nav;
