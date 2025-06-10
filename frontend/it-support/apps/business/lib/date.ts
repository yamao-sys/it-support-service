export const getDateString = (date: Date): string => {
  return date
    .toLocaleDateString("ja-jp", { year: "numeric", month: "2-digit", day: "2-digit" })
    .replaceAll("/", "-");
};
