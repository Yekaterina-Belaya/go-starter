export interface IVacancy {
    id: string;
    title: string;
    description: string;
    location: string;
    salary: number;
    companyId: string;
    createdAt: Date;
    updatedAt: Date;
}