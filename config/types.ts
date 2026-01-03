// types.ts
export type DataParserConfig = {
  inputPath: string;
  outputPath: string;
  fields: {
    [key: string]: {
      type: 'string' | 'number' | 'boolean';
      format?: 'date' | 'time';
    };
  };
};

export type DataParserResult = {
  parsedData: any[];
  errors: string[];
};

export type FieldConfig = {
  type: 'string' | 'number' | 'boolean';
  format?: 'date' | 'time';
};