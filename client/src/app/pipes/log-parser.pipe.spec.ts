import { LogParserPipe } from './log-parser.pipe';

describe('LogParserPipe', () => {
  it('create an instance', () => {
    const pipe = new LogParserPipe();
    expect(pipe).toBeTruthy();
  });
});
