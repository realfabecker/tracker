import dayjs from "dayjs";

export const asDate = (d: string) => {
  return dayjs(d).format("DD/MM/YYYY");
};

export const asBrl = (d: number) => {
  return new Intl.NumberFormat("pt-BR", {
    style: "currency",
    currency: "BRL",
  }).format(d);
};
