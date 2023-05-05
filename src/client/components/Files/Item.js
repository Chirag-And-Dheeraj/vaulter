import Image from "next/image";

const Item = ({ name, lastModified, size, icon }) => {
    return (
        <>
            <section className="grid grid-cols-12 items-center p-2 md:px-2 md:py-3">
                <Image
                    className="col-start-1 fill-black"
                    src={`/icons/${icon}.svg`}
                    alt="Files Icon"
                    width={24}
                    height={24}
                ></Image>
                <div className="col-start-3 col-end-9 sm:col-end-5 md:col-start-2 md:col-end-5">
                    <h3>{name}</h3>
                    <small className="text-xs sm:hidden">
                        Modified {lastModified}
                    </small>
                </div>
                <div className="hidden sm:block sm:col-start-7 sm:col-end-9 md:justify-self-start">
                    {lastModified}
                </div>
                <div className="hidden sm:block col-start-10 col-end-12 md:justify-self-start">
                    {size}
                </div>
                <button className="col-start-12 col-end-13 md:justify-self-end">
                    <Image
                        src="/icons/dots.svg"
                        width={24}
                        height={24}
                        alt="3 Dots Icon"
                    ></Image>
                </button>
            </section>
            <hr className="hidden md:block" />
        </>
    );
};

export default Item;
