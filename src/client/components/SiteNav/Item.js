import Image from "next/image";

const Item = ({ src, alt, name }) => {
    return (
        <button className="flex flex-col place-items-center md:flex-row">
            <Image
                src={src}
                alt={alt}
                width={24}
                height={24}
            ></Image>
            <span className="md:m-2">{name}</span>
        </button>
    );
};

export default Item;
