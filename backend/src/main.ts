import { NestFactory } from '@nestjs/core';
import { DocumentBuilder, SwaggerModule } from '@nestjs/swagger';
import { AppModule } from './app.module';

async function bootstrap() {
  const app = await NestFactory.create(AppModule);
  const config = new DocumentBuilder()
    .setTitle('Blog-project') //the title you want for your swagger docs
    .setDescription('Pet project blog API description') //description
    .setVersion('1.0')  //version setting for the docs
    .build();

  const document = SwaggerModule.createDocument(app, config, {
    include: []
  })
  SwaggerModule.setup('/docs', app, document)
  await app.listen(8001);
}
bootstrap();
